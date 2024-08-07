from .entities import Product
from .serializers import ProductSerializers, UserSerializers
from rest_framework.generics import GenericAPIView
from rest_framework.mixins import (
    ListModelMixin,
    CreateModelMixin,
    RetrieveModelMixin,
    UpdateModelMixin,
    DestroyModelMixin
)
from .models import Product
from rest_framework.pagination import PageNumberPagination


class ProductList(ListModelMixin, CreateModelMixin, GenericAPIView):
    """Представление для получения списков товаров и создание нового товара."""
    serializer_class = ProductSerializers

    def get_queryset(self):
        quaryset = Product.objects.all()
        item_name = self.request.query_params.get('name')
        if item_name:
            quaryset = quaryset.filter(name=item_name)
        return quaryset

    def get(self, request):
        return self.list(request)

    def post(self, request):
        return self.create(request)


class ProductDetail(UpdateModelMixin, RetrieveModelMixin, DestroyModelMixin, GenericAPIView):
    """Представление для получения детальной информации о товаре,
    а также для его редактирования и удаления"""

    queryset = Product.objects.all()
    serializer_class = ProductSerializers

    def get(self, request, *args, **kwargs):
        return self.retrieve(request, *args, **kwargs)

    def put(self, request, *args, **kwargs):
        return self.update(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        return self.destroy(request, *args, **kwargs)


class Paginatorlist(PageNumberPagination):
    page_size = 5



