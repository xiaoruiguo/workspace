from django.contrib import admin
from polls.models import Poll

# Register your models here.

class PollAdmin(admin.ModelAdmin):
    fieldsets = [
            ('XXXXXX', {'fields':['question']}),
            ('Date information xxx', {'fields':['pub_date']}),
        ]

admin.site.register(Poll, PollAdmin)
