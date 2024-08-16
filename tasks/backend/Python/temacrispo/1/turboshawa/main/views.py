from django.views.generic import TemplateView
from heavenlymenu.models import Category


class IndexView(TemplateView):
    template_name = 'main/index.html'

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['categories'] = Category.objects.all()
        context['pagename'] = 'Турбошаверма'
        return context


class AboutView(TemplateView):
    template_name = 'main/about.html'

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['pagename'] = 'О Турбошаверме'
        context['content'] = "Расскажем о себе"
        context['text_about_us'] = ("Мы делаем самую странную и иногда не самую адекватную шаверму "
                                    "в городе Санкт-Петербурге, вы никогда не знаете,что ваш ждем, "
                                    "когда вы придете к нам, это будет сюрприз, буквально!")
        return context
