from django.shortcuts import render
from rest_framework.viewsets import ModelViewSet

from stuff.models import Beer, Shaurma

from .paginators import StuffPagination
from .serializers import BeerSerializer, ShaurmaSerializer


class StuffViewSet(ModelViewSet):
    """Базовый класс для товаров."""
    pagination_class = StuffPagination


class BeerViewSet(StuffViewSet):
    queryset = Beer.objects.all()
    http_method_names = ['get', ]
    serializer_class = BeerSerializer


class ShaurmaViewSet(StuffViewSet):
    queryset = Shaurma.objects.all()
    http_method_names = ['get', ]
    serializer_class = ShaurmaSerializer
