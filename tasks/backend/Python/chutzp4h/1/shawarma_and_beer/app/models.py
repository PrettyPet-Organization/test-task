from django.db import models

# Create your models here.


class BaseMaterial(models.Model):
    # unique=False, чтобы учитывать одноименные единицы сырья с разным сроком годности
    name = models.CharField(max_length=100, unique=False)
    expiration_date = models.DateTimeField(null=True, blank=True)

    class Meta:
        abstract = True

    def __str__(self):
        return self.name


class IngredientCountable(BaseMaterial):
    quantity = models.PositiveSmallIntegerField()

    class Meta:
        verbose_name = 'Ingredient Countable'
        verbose_name_plural = 'Ingredients Countable'


class IngredientUncountable(BaseMaterial):
    weight = models.FloatField()

    class Meta:
        verbose_name = 'Ingredient Unountable'
        verbose_name_plural = 'Ingredients Unountable'


class DrinkCountable(BaseMaterial):
    volume = models.FloatField()
    quantity = models.PositiveSmallIntegerField()

    class Meta:
        verbose_name = 'Drink Countable'
        verbose_name_plural = 'Drinks Countable'


class DrinkUncountable(BaseMaterial):
    volume = models.FloatField()

    class Meta:
        verbose_name = 'Drink Uncountable'
        verbose_name_plural = 'Drinks Uncountable'


class Product(models.Model):
    name = models.CharField(max_length=100)
    size = models.CharField(max_length=100, null=True, blank=True)
    price = models.FloatField(null=True, blank=True)
    ready_to_cook = models.BooleanField(default=True)

    countable_ingredient = models.ManyToManyField(IngredientCountable, through='RecipeCountable')
    uncountable_ingredient = models.ManyToManyField(IngredientUncountable, through='RecipeUncountable')

    class Meta:
        verbose_name = 'Product'
        verbose_name_plural = 'Products'

    def __str__(self):
        return self.name


class RecipeCountable(models.Model):
    product = models.ForeignKey(Product, on_delete=models.CASCADE)
    countable_ingredient = models.ForeignKey(IngredientCountable, on_delete=models.CASCADE)
    quantity = models.PositiveSmallIntegerField()


class RecipeUncountable(models.Model):
    product = models.ForeignKey(Product, on_delete=models.CASCADE)
    uncountable_ingredient = models.ForeignKey(IngredientUncountable, on_delete=models.CASCADE)
    weight = models.FloatField()


class Order(models.Model):
    STATUS_CHOICES = (
        ('in_processing', 'в обработке'),
        ('confirmed', 'подтвержден'),
        ('canceled', 'отменен'),
        ('ready', 'готов'),
    )
    status = models.CharField(choices=STATUS_CHOICES, max_length=13)
    time_created = models.DateTimeField(auto_now_add=True)
    product = models.ManyToManyField(Product, through='OrderProduct')


class OrderProduct(models.Model):
    order = models.ForeignKey(Order, on_delete=models.CASCADE)
    product = models.ForeignKey(Product, on_delete=models.CASCADE)
    quantity = models.PositiveIntegerField()
