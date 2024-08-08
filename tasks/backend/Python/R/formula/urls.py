from django.urls import path
from . import views

urlpatterns = [
    path('', views.CalculateView.as_view())
]