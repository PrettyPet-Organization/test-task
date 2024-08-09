from django.urls import path, include
from rest_framework.routers import DefaultRouter

from .views import BeerViewSet, ShaurmaViewSet


API_VERSION = 'v1'

router = DefaultRouter()
router.register('beer/', BeerViewSet)
router.register('shaurma/', ShaurmaViewSet)

urlpatterns = [
    path(f'{API_VERSION}', include(router.urls)),
]