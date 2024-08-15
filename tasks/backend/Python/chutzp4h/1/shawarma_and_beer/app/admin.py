from django.contrib import admin

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


# Register your models here.

admin.site.register(RecipeCountable)
admin.site.register(RecipeUncountable)
admin.site.register(Order)
admin.site.register(OrderProduct)


# Можно было сделать абстрактный класс, где указать повторяющиеся элементы list_display, но мне кажется,
# что явно указанными значениями проще управлять
@admin.register(IngredientCountable)
class CountableIngredientAdmin(admin.ModelAdmin):
    list_display = (
        'name',
        'quantity',
        'expiration_date',
    )


@admin.register(IngredientUncountable)
class UncountableIngredientAdmin(admin.ModelAdmin):
    list_display = (
        'name',
        'weight',
        'expiration_date',
    )


@admin.register(DrinkCountable)
class CountableDrinkAdmin(admin.ModelAdmin):
    list_display = (
        'name',
        'volume',
        'quantity',
        'expiration_date',
    )


@admin.register(DrinkUncountable)
class UncountableDrinkAdmin(admin.ModelAdmin):
    list_display = (
        'name',
        'volume',
        'expiration_date',
    )


@admin.register(Product)
class ProductAdmin(admin.ModelAdmin):
    list_display = (
        'name',
        'size',
        'price',
    )
