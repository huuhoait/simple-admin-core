import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { formatToDateTime } from '/@/utils/dateUtil';


const { t } = useI18n();

export const columns: BasicColumn[] = [
  {
    title: t('example.memberRank.trans'),
    dataIndex: 'trans',
    width: 100,
  },
  {
    title: t('example.memberRank.name'),
    dataIndex: 'name',
    width: 100,
  },
  {
    title: t('example.memberRank.description'),
    dataIndex: 'description',
    width: 100,
  },
  {
    title: t('example.memberRank.remark'),
    dataIndex: 'remark',
    width: 100,
  },
  {
    title: t('example.memberRank.code'),
    dataIndex: 'code',
    width: 100,
  },
  {
    title: t('common.createTime'),
    dataIndex: 'createdAt',
    width: 70,
    customRender: ({ record }) => {
      return formatToDateTime(record.createdAt);
    },
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: t('example.memberRank.name'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'description',
    label: t('example.memberRank.description'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'remark',
    label: t('example.memberRank.remark'),
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
    field: 'trans',
    label: t('example.memberRank.trans'),
    component: 'Input',
    required: true,
  },
  {
    field: 'name',
    label: t('example.memberRank.name'),
    component: 'Input',
    required: true,
  },
  {
    field: 'description',
    label: t('example.memberRank.description'),
    component: 'Input',
    required: true,
  },
  {
    field: 'remark',
    label: t('example.memberRank.remark'),
    component: 'Input',
    required: true,
  },
  {
    field: 'code',
    label: t('example.memberRank.code'),
    component: 'Input',
    required: true,
  },
];
