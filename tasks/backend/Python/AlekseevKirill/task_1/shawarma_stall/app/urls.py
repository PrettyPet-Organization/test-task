from rest_framework.routers import DefaultRouter
from django.urls import path
from app.views import (
    ProductViewSet,
    OrderView,
)

router = DefaultRouter()
router.register('products', ProductViewSet)

urlpatterns = [
    path('orders/', OrderView.as_view()),
] + router.urls