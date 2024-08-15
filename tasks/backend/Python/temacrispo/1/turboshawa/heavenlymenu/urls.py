from django.urls import path
from . import views

app_name = "goods"

urlpatterns = [
    path('<slug:category_slug>/', views.CatalogView.as_view(), name='index'),
    path("product/<slug:product_slug>/", views.LotView.as_view(), name="lot"),
]
