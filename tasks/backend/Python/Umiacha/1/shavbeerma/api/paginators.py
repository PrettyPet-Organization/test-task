from rest_framework.pagination import PageNumberPagination


class StuffPagination(PageNumberPagination):
    """Пагинатор для товаров."""
    page_size_query_param = 'limit'