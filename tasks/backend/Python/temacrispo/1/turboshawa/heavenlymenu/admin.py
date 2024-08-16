from django.contrib import admin

from heavenlymenu.models import SpicyCategory, Category, Good

# Register your models here.

# admin.site.register(Categories)
# admin.site.register(BiddingCateories)
# admin.site.register(Lots)

@admin.register(Category)
class CategoryAdmin(admin.ModelAdmin):
    prepopulated_fields = {'slug': ('name',)}

@admin.register(SpicyCategory)
class SpicyCategoryAdmin(admin.ModelAdmin):
    prepopulated_fields = {'slug': ('spicy',)}

@admin.register(Good)
class GoodAdmin(admin.ModelAdmin):
    prepopulated_fields = {'slug': ('name',)}