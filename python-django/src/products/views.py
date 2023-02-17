from django.shortcuts import render
from .models import Product
from .forms import ProductForm, RawProductForm


# Create your views here.

def product_create_view(request):
    new_form = RawProductForm()
    if request.method == "POST":
        new_form = RawProductForm(request.POST)
        if new_form.is_valid():
            print(new_form.cleaned_data)
            Product.
        else:
            print(new_form.errors)
    context = {
        "form": new_form,
    }
    return render(request, "products/product_create.html", context)


# def product_create_view(request):
#     # print(request.GET)
#     # print(request.POST)
#     if request.method == "POST":
#         new_title = request.POST.get("title")
#         # Product.object.create(title=new_title)
#         print(new_title)
#     context = {}
#     return render(request, "products/product_create.html", context)

# def product_create_view(request):
#     form = ProductForm(request.POST or None)
#     if form.is_valid():
#         form.save()
#         form = ProductForm()
#
#     context = {
#         "form": form
#     }
#     return render(request, "products/product_create.html", context)

def product_detail_view(request):
    obj = Product.objects.get(id=1)

    # Inefficient way to do that:
    # context = {
    #     "title": obj.title,
    #     "price": obj.price,
    #     "description": obj.description,
    # }

    # More efficient way:
    context = {
        "object": obj
    }
    return render(request, "products/product_detail.html", context)