<template lang="pug">
  i(:class="iconclass")
</template>

<script>
export default {
  data: function() {
    return {
      iconset: 'fontawesome',
        iconTypes: {
          fontawesome: {
            name: 'Font Awesome',
            template: 'fa fa-fw fa-',
            icons: {
              ajax: 'circle-o-notch fa-spin',
              menu: 'bars',
              close: 'times',
              search: 'search',
              compose: 'pencil',
              emoji: 'smile-o',
              attachment: 'paperclip',
              send: 'paper-plane',
            }
          },
          materialdesign: {
            name: 'Material Design',
            template: 'mdi mdi-',
            icons: {
              ajax: '',
              menu: 'menu',
              search: 'magnify',
              close: 'close',
              compose: 'pencil',
              emoji: 'emoticon',
              attachment: 'paperclip',
              send: 'send',
            }
          }
        },
      };
    },
    props: {
      icon: {
        type: String,
        required: true,
      },
      classes: {
        type: String,
        required: false,
        default: ''
      },
    },
    methods: {
    },
    computed: {
      iconclass: function() {
        let iconset = this.iconTypes[this.iconset];
        return iconset.template + iconset.icons[this.icon] + this.classes;
      },
    },
    mounted: function() {
      this.iconset = this.$root.$children[0].iconset;
      this.$root.$children[0].icons.push(this);
    }
}
</script>
