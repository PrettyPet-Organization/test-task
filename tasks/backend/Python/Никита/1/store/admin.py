from django.contrib import admin
from .models import ProductModel


@admin.register(ProductModel)
class ProductAdmin(admin.ModelAdmin):
    model = ProductModel
    list_display = ('pk', 'name', 'price')
    list_display_links = ('pk', 'name')
    fields = ('name', 'price', 'image')