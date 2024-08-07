from django.shortcuts import render
from .forms import FormuleForm


def FormuleView(request):
    """Функция View для отображения Формулы."""
    form = FormuleForm()

    if request.method == "POST":
        form = FormuleForm(request.POST)
        a = int(request.POST.get("a"))
        b = int(request.POST.get("b"))
        text = f'({a} + {b} )**2 = ' \
               f'{( ( (a  ** 2) + (((a * b)) * 2) + ((b ** 2))))}'
        context = {"content": text}
        return render(request, 'index.html', context=context)

    context = {'form': form}
    return render(request, 'index.html', context=context)
