from flask import Flask, request, jsonify
from flask_cors import CORS, cross_origin
import requests

# The only dependency is Flask(pip install flask)

app = Flask(__name__, static_folder = '/static')
cors = CORS(app)

@app.route('/api/v1/register', methods = ['POST', 'GET'])
@cross_origin()
def register():
  return '12345'# fake token

@app.route('/api/v1/topics', methods = ['GET'])
@cross_origin()
def get_all_catetory():
  return jsonify(['Java', 'Python', 'Spring', 'Django'])# fake token

@app.route('/data', methods = ['GET'])
@cross_origin()
def get_all_data():
  res = [{
      "id": 111,
      "story": 'aaaaa',
      "team": 'layercake',
      "epic": 'image cropping',
      "compledDate": '2020-01-02',
      "ticketType": 'chore',
      "cycleTime": '4',
    },
    {
      "id": 222,
      "story": 'aaaaa',
      "team": 'layercake',
      "epic": 'image cropping',
      "compledDate": '2020-02-02',
      "ticketType": 'chore',
      "cycleTime": '5',
    }
  ]
  return jsonify(res)

posts = {
  "posts": [{
      "author": "MS1",
      "title": "aws title11111",
      "tagList": ["aa", "bb", "cc", "dd"],
      "commentCount": 12
    },
    {
      "author": "MS2",
      "title": "aws title12222",
      "tagList": ["ee", "ff", "gg", "hh"],
      "commentCount": 14
    },
    {
      "author": "MS3",
      "title": "aws title333",
      "tagList": ["ee", "ff", "gg", "hh"],
      "commentCount": 20
    },
    {
      "author": "MS3",
      "title": "aws title333",
      "tagList": ["ee", "ff", "gg", "hh"],
      "commentCount": 20
    }
  ]
}

@app.route('/api/v1/posts', methods = ['GET'])
@cross_origin()
def get_all_posts():
  return jsonify(posts), 200