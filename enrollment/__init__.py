from flask import Flask

app = Flask(__name__)

# really funny way to set routes be carefull while doing this
# there is a chance if the order is not right the app will keep crashing
from enrollment import routes
