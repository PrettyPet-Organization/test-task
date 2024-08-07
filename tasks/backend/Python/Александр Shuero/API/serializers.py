from django.contrib.auth.models import User
from rest_framework import serializers
from .models import Product


class UserSerializers(serializers.HyperlinkedModelSerializer):
    """Представление списков пользователей"""
    class Meta:
        model = User
        fields = [
            'id', 'is_staff', 'username'
        ]


class ProductSerializers(serializers.ModelSerializer):
    """Представление списков товаров"""
    class Meta:
        model = Product
        fields = ['id', 'name', 'price', 'count']
