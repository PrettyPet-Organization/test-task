from django.db import models
from django.urls import reverse


class Category(models.Model):
    name = models.CharField(max_length=150, unique=True, verbose_name="Название")
    slug = models.SlugField(
        max_length=150, unique=True, blank=True, null=True, verbose_name="URL"
    )

    class Meta:
        db_table = "Category"
        verbose_name = "Категорию"
        verbose_name_plural = "Категории"

    def __str__(self):
        return self.name


class SpicyCategory(models.Model):
    spicy = models.CharField(
        max_length=40, blank=False, null=False, verbose_name="Острота"
    )
    slug = models.SlugField(
        max_length=150, unique=True, blank=True, null=True, verbose_name="URL"
    )

    class Meta:
        db_table = "SpicyCategory"
        verbose_name = "Острота"

    def __str__(self):
        return self.spicy


class Good(models.Model):
    name = models.CharField(max_length=150, unique=True, verbose_name="Наименование")
    slug = models.SlugField(max_length=150, unique=True, blank=True, null=True, verbose_name="URL")
    description = models.TextField(max_length=600, blank=True, null=True, verbose_name="Описание")
    image = models.ImageField(upload_to="lots_images", blank=True, null=True, verbose_name="Изображение")
    price = models.DecimalField(default=0.00, max_digits=20, decimal_places=2, verbose_name="Цена")
    spicy = models.ForeignKey(to=SpicyCategory, on_delete=models.PROTECT, verbose_name="Острота")
    category = models.ForeignKey(to=Category, on_delete=models.CASCADE, verbose_name="Категория")

    class Meta:
        db_table = "goods"
        verbose_name = "Товар"
        verbose_name_plural = "Товары"

    def __str__(self):
        return self.name

    def get_absolute_url(self):
        return reverse("catalog:product", kwargs={"good_slug": self.slug})

    def sell_price(self):
        return self.price

    def display_id(self):
        return f"{self.id}"