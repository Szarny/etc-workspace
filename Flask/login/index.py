from flask import Flask, request, render_template
from auth import valid_credential

app = Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html")


@app.route("/login", methods=["GET", "POST"])
def login():
    error_message = ""

    if request.method == "POST":
        if valid_credential(username=request.form["username"], password=request.form["password"]):
            return "Hello, {}!".format(request.form["username"])
        else:
            error_message = "Login failed."
    
    return render_template("login.html", error_message=error_message)


if __name__ == '__main__':
    app.run(debug=True)