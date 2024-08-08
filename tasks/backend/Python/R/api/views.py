from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from .models import Product, Order
from .serializers import ProductSerializer, OrderSerializer
from .utils import create_order


class ProductListCreateAPIView(APIView):
    def get(self, request, pk=None):
        if pk:
            try:
                product = Product.objects.get(pk=pk)
            except Product.DoesNotExist:
                return Response({"error": "Product not found."}, status=status.HTTP_404_NOT_FOUND)
            serializer = ProductSerializer(product)
            return Response({"product": serializer.data}, status=status.HTTP_200_OK)

        products = Product.objects.all()
        serializer = ProductSerializer(products, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

    def post(self, request):
        serializer = ProductSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    def patch(self, request, pk):
        try:
            product = Product.objects.get(pk=pk)
        except Product.DoesNotExist:
            return Response({"error": "Product not found."}, status=status.HTTP_404_NOT_FOUND)

        serializer = ProductSerializer(product, data=request.data, partial=True)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_200_OK)

        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class OrderListCreateAPIView(APIView):
    def get(self, request):
        orders = Order.objects.all()
        serializer = OrderSerializer(orders, many=True)
        return Response({"orders": serializer.data}, status=status.HTTP_200_OK)

    def post(self, request):
        serializer = OrderSerializer(data=request.data)
        if serializer.is_valid():
            try:
                order = create_order(serializer=serializer)
                return Response({"order": OrderSerializer(order).data}, status=status.HTTP_201_CREATED)
            except ValueError as e:
                return Response({"error": str(e)}, status=status.HTTP_400_BAD_REQUEST)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
