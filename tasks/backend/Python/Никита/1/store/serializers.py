from rest_framework.serializers import ModelSerializer
from .models import ProductModel


class ProductSerializer(ModelSerializer):
    class Meta:
        model = ProductModel
        fields = ('id', 'name', 'price', 'image')