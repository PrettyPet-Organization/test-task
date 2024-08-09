from django.contrib import admin
from django.urls import path, include

# TODO: сгенерировать документацию и прикрутить её в режиме дебага.
urlpatterns = [
    path('admin/', admin.site.urls),
    # path('api/', include('api.urls')),  # TODO: составить эндпоинты для API.
]

# if settings.DEBUG:  # Вариант раздачи статики для дебага.
#     urlpatterns += static(
#         settings.MEDIA_URL,
#         document_root=settings.MEDIA_ROOT)
#     urlpatterns += static(
#         settings.STATIC_URL,
#         document_root=settings.STATIC_ROOT)
