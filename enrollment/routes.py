from enrollment import app
from flask import render_template
from flask import request
from flask import Response
from json import load, loads, dump, dumps

usersDict = {
    "admin":{"secret":"admin", "fullName":"Administrator of enrollment", "permissions":"all"}
}

coursesDict = {
"1101":{"courseTitle":"Advanced Java", "credits":30, "duration":"20 hours"},
"1102":{"courseTitle":"Advanced C++", "credits":40, "duration":"20 hours"},
"1103":{"courseTitle":"Java Fundamentals", "credits":10, "duration":"10 hours"},
"1104":{"courseTitle":"Algorithms and Data Structures in Java", "credits":50, "duration":"50 hours"},
}

@app.route("/")
@app.route("/index")
def index():
    return render_template("index.html", index=True)

@app.route("/login")
def login():
    return render_template("login.html", login=True)

@app.route("/courses")
def courses():
    return render_template("courses.html", coursesDict=coursesDict, courses=True)

@app.route("/register")
def register():
    return render_template("register.html", register=True)

@app.route("/enroll", methods=["GET", "POST"])
def enroll():
    data = None
    if request.method == "POST":
        data = {"courseID":request.form.get('courseID'),
        "title":request.form.get("title"),
        "duration":request.form.get("duration")}
    else:
        data = {"courseID":request.args.get('courseID'),
        "title":request.args.get("title"),
        "duration":request.args.get("duration")}

    return render_template('enrollment.html', data=data)

@app.route("/api/")
def api():
    return Response(dumps({"apiversion":"1", "description":"enrollment API"}), mimetype="application/json")

@app.route('/api/courses/')
@app.route("/api/courses/<int:idx>")
def apiCourses(idx=None):
    if idx == None:
        return Response(dumps(coursesDict), mimetype="application/json")
    else:
        coursesKeys = list(coursesDict)
        if idx >= len(coursesKeys):
            return Response(dumps({"error":"Course not found"}), mimetype="application/json")
        return Response(dumps(coursesDict[list(coursesDict)[idx]]), mimetype="application/json")

@app.route('/users/')
@app.route('/users/<path:opname>', methods=["POST"])
def users(opname="default"):
    if opname == "default":
        return json.dumps(usersDict)

 