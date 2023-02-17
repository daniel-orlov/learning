from django.shortcuts import render
from django.http import HttpResponse
# Create your views here.
def home_view(request, *args, **kwargs):
    # return HttpResponse("<h1>Hello World, that's me again</h1>") # a string of HTML code
    return render(request, "home.html", {})

def contact_view(request, *args, **kwargs):
    my_context = {
        "username": request.user,
        "page_num": 1,
        "nums": [3, 14, 15, 92, 65]
    }
    return render(request, "contact.html", my_context)