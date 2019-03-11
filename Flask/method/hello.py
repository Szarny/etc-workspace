from flask import Flask, request

app = Flask(__name__)

@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "GET":
        return "Login page"
    elif request.method == "POST":
        return "Login API"