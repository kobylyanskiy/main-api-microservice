import json
import requests
from flask import Flask, request


app = Flask(__name__)


@app.route('/agents', methods=['GET', 'POST'])
def agents():
    if request.method == 'POST':
        agent = request.get_json(force=True)
        #response_neo4j = requests.post('http://neo4j-api:5000/nodes', data=operation).json()        
        response = requests.post('http://mongo-api:5000/agents', data=json.dumps(agent)).json()
        return json.dumps(response)
    else:
        response = requests.get('http://mongo-api:5000/agents').json()
        return json.dumps(response)


@app.route('/agents/<string:codename>', methods=['GET', 'POST'])
def agent_req(codename):
    if request.method == 'POST':
        agent = request.get_json(force=True)
        response = requests.post('http://mongo-api:5000/agents/{}'.format(codename), data=json.dumps(agent)).json()
        return json.dumps(response)
    else:
        response = requests.get('http://mongo-api:5000/agents/{}'.format(codename)).json()
        return json.dumps(response)


@app.route('/operations', methods=['GET', 'POST'])
def operations():
    if request.method == 'POST':
        operation = request.get_json(force=True)
        #check agents(mongo)
        response_neo4j = {'result':'kik'}
        response = requests.post('http://0.0.0.0:5000/operations', data=json.dumps(operation)).json()
        #response_neo4j = requests.post('http://neo4j-api:5000/nodes', data=operation).json()
        #make mongo update
        return json.dumps({
            'result_cassandra': response['result'],
            'result_neo4j': response_neo4j['result'],
        })
    else:
        response = requests.get('http://0.0.0.0:6000/operations').json()
        return json.dumps(response)



@app.route('/operations/<string:codename>', methods=['GET', 'POST'])
def operation_req(codename):
    if request.method == 'POST':
        operation = request.get_json(force=True)
        response = requests.post('http://0.0.0.0:5000/operations/{}'.format(codename), data=json.dumps(operation)).json()
        #if response['result'] == True and response['data']['status'] == ('complete' or 'failed'):
            #for codename in response['data']['agents']:
                ##if response['data']['status'] == 'complete':
                #    rank = 1
                #else:
                #    rank = -1
               # agent = {'rank': rank*response['data']['difficulty']}
                #response_mongo = requests.post('http://mongo-api:5000/agents/{}'.format(codename), data=agent).json()            
        return json.dumps({
            'result': response['result'], 
        })
    else:
        response = requests.get('http://0.0.0.0:6000/operations/{}'.format(codename)).json()
        return json.dumps(response) 



if __name__ == '__main__':
    app.run(host='0.0.0.0')


