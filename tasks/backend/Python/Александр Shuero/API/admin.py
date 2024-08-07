from django.contrib import admin
from .models import Product
from django import forms


class ProductAdminForm(forms.ModelForm):
    """Форма для модели VK"""
    class Meta:
        model = Product
        fields = "__all__"


@admin.register(Product)
class ProductAdmin(admin.ModelAdmin):
    """Админ панель модель VK"""

    list_display = [
        "name",
        "price",
        "count",
    ]
    form = ProductAdminForm
