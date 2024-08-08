from rest_framework.views import APIView
from rest_framework.viewsets import GenericViewSet
from rest_framework import status
from rest_framework.response import Response
from rest_framework.mixins import (
    ListModelMixin,
    RetrieveModelMixin,
    CreateModelMixin
)

from app.models import (
    Product,
    Order
)
from app.serializers import (
    ProductSerializer,
    OrderSerializer,
)


class ProductViewSet(ListModelMixin, CreateModelMixin, RetrieveModelMixin, GenericViewSet):
    """Представление для получения списка продуктов и продукта по id"""

    serializer_class = ProductSerializer
    queryset = Product.objects.all()


class OrderView(APIView):
    """Представление для получения списка заказов и создания заказа"""
    
    def get(self, request):
        orders = Order.objects.all()
        order_serializer = OrderSerializer(orders, many=True)
        return Response({"orders": order_serializer.data}, status=status.HTTP_200_OK)
    
    def post(self, request):
        order_serializer = OrderSerializer(data=request.data)
        if order_serializer.is_valid():
            try:
                order = order_serializer.save()
                return Response({"order": OrderSerializer(order).data}, status=status.HTTP_201_CREATED)
            except ValueError as err:
                return Response({"error": str(err)}, status=status.HTTP_400_BAD_REQUEST)
        return Response(order_serializer.errors, status=status.HTTP_400_BAD_REQUEST)
