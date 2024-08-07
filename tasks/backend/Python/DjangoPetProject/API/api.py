from rest_framework import viewsets
from django.contrib.auth.models import User
from .serializers import UserSerializers, ProductSerializers
from .models import Product


class UserViewSet(viewsets.ModelViewSet):
    """Базовая модель View для отображения пользователей

    """
    queryset = User.objects.all()
    serializer_class = UserSerializers


class ProductViewSet(viewsets.ModelViewSet):
    """Базовая модель View для отображения продуктов

    """
    queryset = Product.objects.all()
    serializer_class = ProductSerializers
