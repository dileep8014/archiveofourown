import NProgress from 'nprogress';
import 'nprogress/nprogress.css';

const ProgressOpt = (props: () => void) => {
  NProgress.start();
  props();
  NProgress.done();
};

export default ProgressOpt;
