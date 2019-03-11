from flask import Flask, url_for

app = Flask(__name__)

@app.route("/")
def index():
    return "index"


@app.route("/login")
def login():
    return "login"


@app.route("/profile/<int:id>")
def profile(id):
    return "{}'s profile".format(id)

with app.test_request_context():
    print(url_for("index"))
    print(url_for("login"))
    print(url_for("login", next="/profile"))
    print(url_for("profile", id=123, name="John"))