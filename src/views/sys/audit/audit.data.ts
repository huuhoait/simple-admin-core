import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { formatToDateTime } from '/@/utils/dateUtil';


const { t } = useI18n();

export const columns: BasicColumn[] = [
  {
    title: t('sys.audit.objectName'),
    dataIndex: 'objectName',
    width: 100,
  },
  {
    title: t('sys.audit.actionName'),
    dataIndex: 'actionName',
    width: 100,
  },
  {
    title: t('sys.audit.changedData'),
    dataIndex: 'changedData',
    width: 100,
  },
  {
    title: t('sys.audit.createdBy'),
    dataIndex: 'createdBy',
    width: 100,
  },

];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: t('sys.audit.name'),
    component: 'Input',
    colProps: { span: 8 },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'ID',
    component: 'Input',
    show: false,
  },

  {
    field: 'objectName',
    label: t('sys.audit.objectName'),
    component: 'Input',
    required: true,
  },
  {
    field: 'actionName',
    label: t('sys.audit.actionName'),
    component: 'Input',
    required: true,
  },
  {
    field: 'changedData',
    label: t('sys.audit.changedData'),
    component: 'Input',
    required: true,
  },
  {
    field: 'createdBy',
    label: t('sys.audit.createdBy'),
    component: 'Input',
    required: true,
  },
];
