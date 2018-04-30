import json
from flask import Flask, request


app = Flask(__name__)


@app.route('/agents', methods=['GET', 'POST'])
def agents():
    if request.method == 'POST':
        return json.dumps({
            'result': True,
        })
    else:
        return json.dumps({
            'result': True,
            'data': [{'codename': 'agent'}],
        })


if __name__ == '__main__':
    app.run(host='0.0.0.0')
