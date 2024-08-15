from rest_framework import viewsets

from .models import (
    IngredientCountable,
    IngredientUncountable,

    DrinkCountable,
    DrinkUncountable,

    Product,
    RecipeCountable,
    RecipeUncountable,

    Order,
    OrderProduct,
)
from .serializers import (
    IngredientCountableSerializer,
    IngredientUncountableSerializer,

    DrinkCountableSerializer,
    DrinkUncountableSerializer,

    ProductSerializer,
    RecipeCountableSerializer,
    RecipeUncountableSerializer,

    OrderSerializer,
    OrderProductSerializer,
)

# Create your views here.


class BaseViewSet(viewsets.ModelViewSet):
    model = None

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)

        cls.queryset = cls.model.objects.all()


class IngredientCountableViewSet(BaseViewSet):
    model = IngredientCountable
    serializer_class = IngredientCountableSerializer


class IngredientUncountableViewSet(BaseViewSet):
    model = IngredientUncountable
    serializer_class = IngredientUncountableSerializer


class DrinkCountableViewSet(BaseViewSet):
    model = DrinkCountable
    serializer_class = DrinkCountableSerializer


class DrinkUncountableViewSet(BaseViewSet):
    model = DrinkUncountable
    serializer_class = DrinkUncountableSerializer


class ProductViewSet(BaseViewSet):
    model = Product
    serializer_class = ProductSerializer


class RecipeCountableViewSet(BaseViewSet):
    model = RecipeCountable
    serializer_class = RecipeCountableSerializer


class RecipeUncountableViewSet(BaseViewSet):
    model = RecipeUncountable
    serializer_class = RecipeUncountableSerializer


class OrderViewSet(BaseViewSet):
    model = Order
    serializer_class = OrderSerializer


class OrderProductViewSet(BaseViewSet):
    model = OrderProduct
    serializer_class = OrderProductSerializer
