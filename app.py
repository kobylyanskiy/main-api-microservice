import json
import requests
from flask import Flask, request


app = Flask(__name__)


@app.route('/agents', methods=['GET', 'POST'])
def agents():
    if request.method == 'POST':
        return requests.post('http://mongo-api:5000/agents').content
    else:
        return requests.get('http://mongo-api:5000/agents').content


if __name__ == '__main__':
    app.run(host='0.0.0.0')
