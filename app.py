import json
import requests
from flask import Flask, request


app = Flask(__name__)


@app.route('/agents', methods=['GET', 'POST'])
def agents():
    if request.method == 'POST':
        agent = request.get_json(silent=True)
        response = requests.post('http://mongo-api:5000/agents', data=agent).json()
        return json.dumps({
            'result': response['result'],
        })
    else:
        response = requests.get('http://mongo-api:5000/agents').json()
        if response['result']:
            return json.dumps({
                'result': True,
                'data': [],
                'codefresh': 'test',
            })
        else:
            return json.dumps({
                'result': False,
                'data': [],
                'codefresh': 'test',
            })


@app.route('/agents/<string:codename>', methods=['GET', 'POST'])
def get_agent(codename):
    if request.method == 'POST':
        agent = request.get_json(silent=True)
        response = requests.post('http://mongo-api:5000/agents/{}'.format(codename), data=agent).json()
        return json.dumps({'result': response['result']})
    else:
        response = requests.get('http://mongo-api:5000/agents/{}'.format(codename)).json()
        if not response['result']:
            return json.dumps({
                'result': False,
            })
        return json.dumps({
            'result': True,
            'agent': response['agent'],
        })


if __name__ == '__main__':
    app.run(host='0.0.0.0')
