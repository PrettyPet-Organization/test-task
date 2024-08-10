from rest_framework import serializers
from django.core.validators import MinValueValidator

from decimal import Decimal

from .models import Good


class GoodsSerializer(serializers.HyperlinkedModelSerializer):
    stock = serializers.IntegerField(validators=[MinValueValidator(0)])
    price = serializers.DecimalField(max_digits=7, decimal_places=2, validators=[MinValueValidator(Decimal('0.00'))])

    class Meta:
        model = Good
        fields = ['id', 'name', 'price', 'stock', 'slug', 'description']
