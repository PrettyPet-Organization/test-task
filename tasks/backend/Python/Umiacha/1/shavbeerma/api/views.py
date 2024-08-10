from django.shortcuts import render
from rest_framework.viewsets import ModelViewSet

from stuff.models import Beer, Shaurma

from .paginators import StuffPagination

# TODO: ??Написать сериализаторы для API
class StuffViewSet(ModelViewSet):
    """Базовый класс для товаров."""
    pagination_class = StuffPagination


class BeerViewSet(StuffViewSet):
    queryset = Beer.objects.all()
    http_method_names = ['get', ]


class ShaurmaViewSet(StuffViewSet):
    queryset = Shaurma.objects.all()
    http_method_names = ['get', ]
