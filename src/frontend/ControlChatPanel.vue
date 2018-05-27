 <template lang="pug">
  .chat-panel(:class="{'expanded': isExpanded}")
    .header(@click="expand")
      button.menu.wobbler(@click="$root.$children[0].wobble")
        control-icon(:icon="'menu'")
      span.title {{ title }}
    .search
      input(type="text", placeholder="Search...", v-model="searchInput")
      control-icon.icon(:icon="'search'")
    .content(:ref="'content'")
      control-ajax(v-if="isLoaded === false", :ref="'ajax'")
      ul.user-list
        chat-user(v-for="(user, index) in filteredUsers", :key="index" :userdata="user", :index="index")
    button.compose.wobbler(@click="$root.$children[0].wobble")
      control-icon.icon(:icon="'compose'")
</template>

<script>
import ControlIcon from './ControlIcon.vue'
import ControlAjax from './ControlAjax.vue'
import ControlUser from './ControlChatUser.vue'
import PerfectScrollbar from 'perfect-scrollbar'

export default {
  components: {
    'control-icon': ControlIcon,
    'control-ajax': ControlAjax,
    'chat-user': ControlUser
  },
  data: function() {
    return {
      isExpanded: true,
      isLoaded: false,
      users: [],
      filteredUsers: [],
      searchInput: '',
      scrollbar: false,
    };
  },
  props: {
    title: {
      type: String,
      required: false,
      default: 'Hello World'
    },
  },
  computed: {
  },
  methods: {
    expand: function() {
      this.isExpanded = !this.isExpanded;
    },
    addUser: function(userData) {
      this.users.push(userData)
    },
    filterByName: function(name) {
      let filteredList = [];

      this.users.forEach(function(user) {
        if (user.username.toLowerCase().indexOf(name) !== -1) {
          filteredList.push(user);
        }
      });
      return filteredList;
    },
    hideAjax: function() {
      this.isLoaded = true;
    },
    showAjax: function() {
      this.isLoaded = false;
    },
  },
  watch: {
    searchInput: function(value) {
      if (value.length > 0) {
        this.filteredUsers = this.filterByName(value.toLowerCase());
      } else {
        this.filteredUsers = this.users;
      }
    }
  },
  mounted: function() {
    this.filteredUsers = this.users;
    this.scrollbar = new PerfectScrollbar(this.$refs.content);
  }
}
</script>
