from rest_framework.serializers import (
    ModelSerializer,
    PrimaryKeyRelatedField
)
from rest_framework.exceptions import ValidationError


from app.models import (
    Product,
    Order,
)


class ProductSerializer(ModelSerializer):
    """Сериализатор модели продукта"""

    class Meta:
        model = Product
        fields = '__all__'


class OrderSerializer(ModelSerializer):
    """Сериализатор модели заказа"""

    product = PrimaryKeyRelatedField(queryset=Product.objects.all(), write_only=True)
    product_detail = ProductSerializer(source='product', read_only=True)

    class Meta:
        model = Order
        fields = ['id', 'product', 'product_detail', 'status', 'telephone', 'count']

    def create(self, validated_data):
        validated_data['status'] = 'in_processing'
        return super().create(validated_data)

