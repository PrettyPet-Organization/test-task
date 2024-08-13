from django.shortcuts import render
from rest_framework.viewsets import ModelViewSet

from stuff.models import Beer, Shaurma

from .paginators import StuffPagination
from .serializers import BeerSerializer, ShaurmaSerializer


class StuffViewSet(ModelViewSet):
    """Базовый класс для товаров."""
    http_method_names = ['get', ]
    pagination_class = StuffPagination


class BeerViewSet(StuffViewSet):
    queryset = Beer.objects.all()
    serializer_class = BeerSerializer


class ShaurmaViewSet(StuffViewSet):
    queryset = Shaurma.objects.all()
    serializer_class = ShaurmaSerializer
