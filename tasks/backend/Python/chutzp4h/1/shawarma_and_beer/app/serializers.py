from rest_framework import serializers

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


class BaseSerializer(serializers.ModelSerializer):
    model = None

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)

        cls.Meta = type(
            'Meta',
            (),
            {
                'model': cls.model,
                'fields': '__all__'
            }
        )


class IngredientCountableSerializer(BaseSerializer):
    model = IngredientCountable


class IngredientUncountableSerializer(BaseSerializer):
    model = IngredientUncountable


class DrinkCountableSerializer(BaseSerializer):
    model = DrinkCountable


class DrinkUncountableSerializer(BaseSerializer):
    model = DrinkUncountable


class ProductSerializer(BaseSerializer):
    model = Product


class RecipeCountableSerializer(BaseSerializer):
    model = RecipeCountable


class RecipeUncountableSerializer(BaseSerializer):
    model = RecipeUncountable


class OrderSerializer(BaseSerializer):
    model = Order


class OrderProductSerializer(BaseSerializer):
    model = OrderProduct
