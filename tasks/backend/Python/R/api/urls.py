# shawarma_shop/urls.py
from django.urls import path
from api.views import ProductListCreateAPIView, OrderListCreateAPIView

urlpatterns = [
    path('products/', ProductListCreateAPIView.as_view()),
    path('products/<int:pk>/', ProductListCreateAPIView.as_view()),
    path('orders/', OrderListCreateAPIView.as_view()),
]
