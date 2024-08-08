from django.shortcuts import render
from django.views import View
from .forms import FormulaForm

calc_funcs = {
    'add': lambda a, b: a + b,
    'rm': lambda a, b: a - b,
    'mul': lambda a, b: a * b,
    'div': lambda a, b: a / b,
    'mod': lambda a, b: a % b,
}


class CalculateView(View):
    def get(self, request):
        form = FormulaForm()
        return render(request, 'formula/index.html', {'form': form})

    def post(self, request):
        form = FormulaForm(request.POST)
        if form.is_valid():
            op = form.cleaned_data['operator']
            a = form.cleaned_data['a']
            b = form.cleaned_data['b']
            try:
                result = calc_funcs[op](a, b)
            except ZeroDivisionError:
                result = "Ошибка: нельзя делить на ноль!"
            return render(request, 'formula/index.html', {'form': form, 'result': result})
