from rest_framework import serializers
from .models import Product, Order


class ProductSerializer(serializers.ModelSerializer):
    class Meta:
        model = Product
        fields = '__all__'


class OrderSerializer(serializers.ModelSerializer):
    class Meta:
        model = Order
        fields = '__all__'

    def validate(self, data):
        product = data['product']
        order_quantity = data['quantity']
        if product.quantity < order_quantity:
            raise serializers.ValidationError("Недостаточное количество продукта на складе.")
        return data
