import json
import requests
from flask import Flask, request


app = Flask(__name__)


@app.route('/agents', methods=['GET', 'POST'])
def agents():
    if request.method == 'POST':
        agent = request.get_json(silent=True)
        response = requests.post('http://mongo-api:5000/agents', data=json.dumps(agent)).json()
        return json.dumps(response)
    else:
        response = requests.get('http://mongo-api:5000/agents').json()
        return json.dumps(response)


@app.route('/agents/<string:codename>', methods=['GET', 'POST'])
def agent_req(codename):
    if request.method == 'POST':
        agent = request.get_json(silent=True)
        response = requests.post('http://mongo-api:5000/agents/{}'.format(codename), data=agent).json()
        return json.dumps({'result': response['result']})
    else:
        response = requests.get('http://mongo-api:5000/agents/{}'.format(codename)).json()
        return json.dumps(response)


@app.route('/operations', methods=['GET', 'POST'])
def operations():
    if request.method == 'POST':
        operation = request.get_json(silent=True)
        response = requests.post('http://cassandra-api:5000/operations', data=operation).json()
        response_neo4j = requests.post('http://neo4j-api:5000/nodes', data=operation).json()
        return json.dumps({
            'result': response['result'],
            'result_neo4j': response_neo4j['result'],
        })
    else:
        response = requests.get('http://cassandra-api:5000/operations').json()
        if response['result']:
            return json.dumps({
                'result': True,
                'data': [],
            })
        else:
            return json.dumps({
                'result': False,
                'data': [],
            })


if __name__ == '__main__':
    app.run(host='0.0.0.0')
