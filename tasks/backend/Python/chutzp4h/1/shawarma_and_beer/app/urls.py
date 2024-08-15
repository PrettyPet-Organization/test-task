from django.urls import path, include
from rest_framework.routers import DefaultRouter

from .views import (
    IngredientCountableViewSet,
    IngredientUncountableViewSet,

    DrinkCountableViewSet,
    DrinkUncountableViewSet,

    ProductViewSet,
    RecipeCountableViewSet,
    RecipeUncountableViewSet,

    OrderViewSet,
    OrderProductViewSet,
)


router = DefaultRouter()

router.register(r'ingredients-countable', IngredientCountableViewSet)
router.register(r'ingredients-uncountable', IngredientUncountableViewSet)

router.register(r'drinks-countable', DrinkCountableViewSet)
router.register(r'drinks-uncountable', DrinkUncountableViewSet)

router.register(r'products', ProductViewSet)
router.register(r'recipes-countable', RecipeCountableViewSet)
router.register(r'recipes-uncountable', RecipeUncountableViewSet)

router.register(r'orders', OrderViewSet)
router.register(r'orders-products', OrderProductViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
