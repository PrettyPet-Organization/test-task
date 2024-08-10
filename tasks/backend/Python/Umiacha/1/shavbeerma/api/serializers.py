# TODO: написать сериализаторы.
from rest_framework import serializers

from stuff.models import Beer, Shaurma, Ingredient, Recipe


class IngredientSerializer(serializers.ModelSerializer):
    
    class Meta:
        model = Ingredient
        fields = (
            'id', 'name', 'measurement_unit',
        )


class BeerSerializer(serializers.ModelSerializer):
    
    class Meta:
        model = Beer
        fields = (
            'id', 'name', 'sort', 'category', 'alcohol',
            'density', 'producer', 'price',
        )


class ShaurmaSerializer(serializers.ModelSerializer):
    ingredients = IngredientSerializer(many=True)
    
    class Meta:
        model = Shaurma
        fields = (
            'name', 'ingredients',
        )
    
    def to_representation(self, instance):
        ret = super().to_representation(instance)
        for ingredient in ret['ingredients']:
            ingredient['amount'] = Recipe.objects.get(
                shaurma=Shaurma.objects.get(name=ret['name']),
                ingredient=Ingredient.objects.get(id=ingredient['id'])
            ).amount
        return ret