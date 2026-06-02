import {
  Controller,
  Get,
  HttpCode,
  Param,
  Query,
  Req,
  UseGuards,
} from '@nestjs/common';
import {
  ApiBearerAuth,
  ApiOperation,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { CommonResponse } from 'src/common/dto/common-response.dto';
import { ManagersService } from './manager.service';
import { swaggerExampleUrls } from 'src/config/swagger-example-urls';

@ApiTags('Managers')
@Controller('managers/events')
export class ManagersController {
  constructor(private readonly managerService: ManagersService) {}

  @ApiOperation({ summary: 'Retrieve a specific event by ID' })
  @ApiResponse({
    status: 200,
    description: 'The event details.',
    schema: {
      example: {
        code: 0,
        message: 'Success',
        data: {
          id: 'c672b08f-65df-48e8-a46e-389d18d87eda',
          title: '등록 테스트',
          thumbnail: swaggerExampleUrls.eventLogo,
          place: '카이스트 문지캠퍼스 강의동 407호',
          cast: '정글 9기',
          ageLimit: 12,
          svg: '{"svgString".../svg>"}',
          ticketingStartTime: '2024-12-12T21:16:00.000Z',
          createdAt: '2024-12-05T21:17:00.644Z',
          updatedAt: '2024-12-05T21:17:00.644Z',
          user: {
            id: '3cf8be9f-0b87-4bad-9199-892d6d53be16',
            nickname: '용맹한코랄양치기',
            email: 'manager@jungle.com',
            profileImage: null,
            role: 'manager',
          },
          areas: [
            {
              id: 'f2a7a545-704a-419e-a24b-09eef9753cf9',
              label: 'A',
              price: 50000,
              svg: '<g id="50">...</g>',
              seats: [
                {
                  id: '2df196b8-bb73-4a86-92bb-639a9dc3066c',
                  cx: 734,
                  cy: 836,
                  row: 20,
                  number: 2,
                  reservations: [
                    {
                      id: '733b1a7a-28cb-4ec3-8d07-a2a9e219d2b1',
                      order: {
                        id: '27da7533-92e4-42a6-b14d-e656238f53e7',
                        createdAt: '2024-12-06T07:05:16.026Z',
                        updatedAt: '2024-12-06T07:05:16.026Z',
                        deletedAt: null,
                        user: {
                          id: '8454aee1-adec-4a4a-be2d-38379db653d2',
                          nickname: '지혜로운옥색코끼리',
                          email: '최우석@jungle.com',
                          role: 'user',
                        },
                      },
                    },
                  ],
                },
                {
                  id: '599ff1e3-2edc-4f3f-8fbe-5501ff254a31',
                  cx: 417,
                  cy: 72,
                  row: 1,
                  number: 1,
                  reservations: [],
                },
                {
                  id: 'a87e8a51-b530-4555-9410-0e8ffaa91a71',
                  cx: 318,
                  cy: 836,
                  row: 20,
                  number: 1,
                  reservations: [],
                },
              ],
            },
          ],
        },
      },
    },
  })
  @ApiResponse({
    status: 404,
    description: 'Event not found.',
    schema: {
      example: {
        code: 9,
        message: 'Event not found',
      },
    },
  })
  @ApiResponse({
    status: 500,
    description: 'Internal server error',
    schema: {
      example: {
        code: 6,
        message: 'Internal server error',
      },
    },
  })
  @ApiBearerAuth()
  @Get(':eventId/reservation-status')
  @HttpCode(200)
  @UseGuards(JwtAuthGuard)
  findOneEventWithSeatStatus(
    @Param('eventId') eventId: string,
    @Query('eventDateId') eventDateId: string,
    @Req() req: any,
  ): Promise<CommonResponse<any>> {
    const { userId } = req.user;
    return this.managerService.findOneEventWithSeatStatus(
      userId,
      eventId,
      eventDateId,
    );
  }

  @ApiOperation({ summary: 'Retrieve all events' })
  @ApiResponse({
    status: 200,
    description: 'List of all events.',
    schema: {
      example: {
        code: 0,
        message: 'Success',
        data: [
          {
            id: 'ba1cdf2b-69ec-473f-a501-47a7b1e73602',
            title: 'Music Festival',
            thumbnail: swaggerExampleUrls.eventThumbnail,
            place: 'Central Park',
            cast: 'Famous Band',
            ageLimit: 18,
            ticketingStartTime: '2024-11-23T19:00:00.000Z',
            createdAt: '2024-11-14T23:22:01.274Z',
            updatedAt: '2024-11-14T23:22:01.274Z',
            user: {
              id: 'f55235ae-4496-4982-9b0d-24ec3a61d438',
              nickname: 'new_nickname',
              email: 'se1620236@naver.com',
              profileImage: null,
              role: 'user',
            },
            eventDates: [
              {
                id: '016d3b2c-31b3-408a-a8df-a77b8777b0a3',
                date: '2024-12-01T19:00:00.000Z',
                createdAt: '2024-11-14T23:22:01.274Z',
                updatedAt: '2024-11-14T23:22:01.274Z',
              },
              {
                id: 'aad8e60c-7973-40c6-b53e-335ac03a2af0',
                date: '2024-12-02T19:00:00.000Z',
                createdAt: '2024-11-14T23:22:01.274Z',
                updatedAt: '2024-11-14T23:22:01.274Z',
              },
            ],
          },
          {
            id: '0bf43d9c-4a03-4c8c-a972-759dc1f19a41',
            title: 'Music Festival',
            thumbnail: swaggerExampleUrls.eventThumbnail,
            place: 'Central Park',
            cast: 'Famous Band',
            ageLimit: 18,
            ticketingStartTime: '2024-11-23T19:00:00.000Z',
            createdAt: '2024-11-14T23:22:24.170Z',
            updatedAt: '2024-11-14T23:22:24.170Z',
            user: {
              id: 'f55235ae-4496-4982-9b0d-24ec3a61d438',
              nickname: 'new_nickname',
              email: 'se1620236@naver.com',
              profileImage: null,
              role: 'user',
            },
            eventDates: [
              {
                id: 'bf8e2326-65d1-47d4-b3ca-cda18a741562',
                date: '2024-12-01T19:00:00.000Z',
                createdAt: '2024-11-14T23:22:24.170Z',
                updatedAt: '2024-11-14T23:22:24.170Z',
              },
              {
                id: 'ec7fe4ec-4d16-4f11-87a9-d76eb95bb8ca',
                date: '2024-12-02T19:00:00.000Z',
                createdAt: '2024-11-14T23:22:24.170Z',
                updatedAt: '2024-11-14T23:22:24.170Z',
              },
            ],
          },
        ],
      },
    },
  })
  @ApiResponse({
    status: 500,
    description: 'Internal server error',
    schema: {
      example: {
        code: 6,
        message: 'Internal server error',
      },
    },
  })
  @ApiBearerAuth()
  @Get()
  @HttpCode(200)
  @UseGuards(JwtAuthGuard)
  findAll(@Req() req: any): Promise<CommonResponse<any>> {
    const { userId } = req.user;
    return this.managerService.findAll(userId);
  }
}
