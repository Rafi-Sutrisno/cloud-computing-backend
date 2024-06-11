import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'

import io
import tensorflow as tf
from tensorflow import keras
import numpy as np
from PIL import Image
import sys
import urllib.request as urllib
import base64
from flask import Flask, request, jsonify

model = keras.models.load_model("model.h5")

app = Flask(__name__)

def transform_image(pillow_image):
    # Convert grayscale to RGB
    pillow_image = pillow_image.convert('RGB')
    data = np.asarray(pillow_image)
    data = data / 255.0
    data = data[np.newaxis, ...]
    data = tf.image.resize(data, [224, 224])
    return data

def predict(x):
    predictions = model(x)
    predictions = tf.nn.softmax(predictions)
    pred0 = predictions[0]
    label0 = np.argmax(pred0)
    return label0

@app.route('/predict', methods=['POST'])
def predict_image():
    imageName = request.json.get('imageName')
    if not imageName:
        return jsonify({"error": "No imageName provided"}), 400
    
    urlBucket = "https://storage.googleapis.com/example-bucket-test-cc-trw/prediction/" 
    filepath = urlBucket + imageName

    try:
        contents = urllib.urlopen(filepath)
        image_bytes = np.asarray(bytearray(contents.read()), dtype="uint8")
        pillow_img = Image.open(io.BytesIO(image_bytes)).convert('L')
        tensor = transform_image(pillow_img)
        prediction = predict(tensor)
        data = {"prediction": int(prediction)}
        return jsonify(data)
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=int(os.environ.get('PORT', 8000)))