from django.contrib import admin

from .models import Beer, Shaurma, Ingredient

# TODO: добавить админки для моделей.
@admin.register(Beer)
class BeerAdmin(admin.ModelAdmin):
    pass


@admin.register(Shaurma)
class ShaurmaAdmin(admin.ModelAdmin):
    pass


@admin.register(Ingredient)
class IngredientAdmin(admin.ModelAdmin):
    pass
