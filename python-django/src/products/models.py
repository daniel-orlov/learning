from django.db import models

# Create your models here.

class Product(models.Model):
    title = models.CharField(max_length=100) #max_length required
    description = models.TextField(blank=True, null=True)
    price = models.DecimalField(max_digits=7, decimal_places=2)
    summary = models.TextField()
    featured = models.BooleanField(default=True)