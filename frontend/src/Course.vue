<template>
  
<tbody id="course_table_body">
  <tr class="course_header" :class="{'unscheduled_course': !course.scheduled && allowed}">
    <td>
      <input
        type="checkbox"
        v-model="allowed"
      >
    </td>
    <td colspan="2">{{ title }}</td>
  </tr>
  <priority-buttons :course="course" />
  <group
    v-for="group in course.groups"
    :key="group.id"
    :group="group"
    :schedule="schedule"
    @allowedItemsChanged="$emit('allowedItemsChanged')"
  />
</tbody>

</template>
<script>
  
import PriorityButtons from './PriorityButtons.vue'
import GroupComponent from './Group.vue'

export default {
  props: ['course', 'schedule'],
  computed: {
    'allowed': {
      get () {
        return this.course.allowed
      },
      set (value) {
        this.course.allowed = value
        this.$emit('allowedItemsChanged')
      }
    },

    title() {
      return `${this.course.code} ${this.course.name} (${this.course.getTypeName()})`
    }
  },
  methods: {

    groupClass (group) {
      for (const term of this.schedule) {
        if (term.group === group) {
          return 'selected_course'
        }
      }
      return ''
    }

  },
  components: {
    'priority-buttons': PriorityButtons,
    'group': GroupComponent
  }
}

</script>
