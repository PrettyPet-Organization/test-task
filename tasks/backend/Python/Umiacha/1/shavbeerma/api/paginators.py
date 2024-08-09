from rest_framework.pagination import PageNumberPagination


class StuffPagination(PageNumberPagination):
    """Пагинатор для товаров."""
    page_size = 10  # TODO: перенести эту константу в настройки (есчо, можно будет наследовать другой пагинатор от этого и переписать).
    page_size_query_param = 'limit'