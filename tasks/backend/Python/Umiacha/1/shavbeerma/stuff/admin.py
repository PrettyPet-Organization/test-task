from django.contrib import admin

from .models import Beer, Shaurma, Ingredient


class IngredientInline(admin.StackedInline):
    model = Shaurma.ingredients.through
    extra = 0


@admin.register(Beer)
class BeerAdmin(admin.ModelAdmin):
    pass


@admin.register(Shaurma)
class ShaurmaAdmin(admin.ModelAdmin):
    list_display = ('name', 'get_ingredients')
    inlines = (IngredientInline,)
    
    @admin.display(description='Ingredients')
    def get_ingredients(self, obj):
        return obj.ingredients.all()


@admin.register(Ingredient)
class IngredientAdmin(admin.ModelAdmin):
    pass


admin.site.empty_value_display = 'Не задано'
