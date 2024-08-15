from django.views.generic import ListView, DetailView
from django.http import Http404
from heavenlymenu.models import Good


class CatalogView(ListView):
    model = Good
    template_name = "heavenlymenu/catalog.html"
    context_object_name = "goods"
    slug_url_kwarg = "category_slug"
    paginate_by = 10
    allow_empty = False

    def get_queryset(self):
        category_slug = self.kwargs.get(self.slug_url_kwarg)

        if category_slug == "all":
            goods = super().get_queryset()
        else:
            goods = super().get_queryset().filter(category__slug=category_slug)
        return goods

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context["pagename"] = "Турбошаверма"
        context["slug_url"] = self.kwargs.get(self.slug_url_kwarg)
        return context


class LotView(DetailView):
    template_name = "heavenlymenu/lot.html"
    slug_url_kwarg = "product_slug"
    context_object_name = "product"

    def get_object(self, queryset=None):
        product = Good.objects.get(slug=self.kwargs.get(self.slug_url_kwarg))
        return product

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context["pagename"] = self.object.name
        return context
