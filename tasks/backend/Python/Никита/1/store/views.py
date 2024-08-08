from rest_framework.viewsets import GenericViewSet
from rest_framework.mixins import ListModelMixin, RetrieveModelMixin
from drf_spectacular.utils import extend_schema

from .models import ProductModel
from .serializers import ProductSerializer


class ProductViewSet(ListModelMixin, RetrieveModelMixin, GenericViewSet):
    serializer_class = ProductSerializer
    queryset = ProductModel.objects.all()

    @extend_schema(tags=['Products'])
    def list(self, request, *args, **kwargs):
        return super().list(request, *args, **kwargs)
    
    @extend_schema(tags=['Products'])
    def retrieve(self, request, *args, **kwargs):
        return super().retrieve(request, *args, **kwargs)